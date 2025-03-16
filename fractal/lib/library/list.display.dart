import 'package:fixnum/fixnum.dart' as fixnum;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:pasteboard/pasteboard.dart' as pasteboard;
import 'package:desktop_drop/desktop_drop.dart';
import 'package:fractal/designkit.dart' as ds;
import 'package:fractal/media.dart' as media;
import 'package:fractal/mimex.dart' as mimex;
import './search.row.dart';

class AvailableListDisplay extends StatefulWidget {
  final media.FnMediaSearch search;
  final media.FnUploadRequest upload;
  const AvailableListDisplay({
    super.key,
    this.search = media.media.get,
    this.upload = media.media.upload,
  });

  @override
  State<StatefulWidget> createState() => _AvailableListDisplay();
}

class _AvailableListDisplay extends State<AvailableListDisplay> {
  bool _loading = true;
  Widget? _player = null;
  ds.Error? _cause = null;
  media.MediaSearchResponse _res = media.media.response(
    next: media.media.request(limit: 32),
  );

  void refresh() {
    widget
        .search(_res.next)
        .then((v) {
          setState(() {
            _res = v;
            _loading = false;
          });
        })
        .catchError((e) {
          setState(() {
            _cause = ds.Error.unknown(e);
            _loading = false;
          });
        });
  }

  @override
  void initState() {
    super.initState();
    refresh();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final upload = (DropDoneDetails v) {
      setState(() {
        _loading = true;
      });

      final multiparts = v.files.map((c) {
        return media.media.uploadable(c.path, c.name, c.mimeType!);
      });

      return Future.microtask(() {
        return Future.wait(
              multiparts.map((fv) {
                return fv.then((v) {
                  return widget
                      .upload((req) {
                        req..files.add(v);
                        return req;
                      })
                      .then((uploaded) {
                        setState(() {
                          _res.items.add(uploaded.media);
                        });
                      })
                      .catchError((cause) {
                        setState(() {
                          _cause = ds.Error.unknown(cause);
                        });
                      });
                });
              }),
            )
            .then((v) => ds.NullWidget)
            .catchError(ds.Error.unknown)
            .whenComplete(
              () => setState(() {
                _loading = false;
              }),
            );
      });
    };

    return CallbackShortcuts(
      bindings: <ShortcutActivator, VoidCallback>{
        const SingleActivator(LogicalKeyboardKey.keyV, control: true): () {
          print("PASTE ACTIVATED");
          pasteboard.Pasteboard.files().then((files) {
            print("DERP DERP ${files.length}");
          });
        },
      },
      child: ds.Overlay(
        overlay: _player,
        child: ds.Table(
          loading: _loading,
          cause: _cause,
          leading: SearchTray(
            onSubmitted: (v) {
              setState(() {
                _res.next.query = v;
                _res.next.offset = fixnum.Int64(0);
              });
              refresh();
            },
            next: (i) {
              setState(() {
                _res.next.offset = i;
              });
              refresh();
            },
            current: _res.next.offset,
            empty: fixnum.Int64(_res.items.length) < _res.next.limit,
            trailing: ds.FileDropWell(
              upload,
              child: IconButton(
                onPressed: () {},
                icon: Icon(Icons.file_upload_outlined),
              ),
            ),
          ),
          children: _res.items,
          (v) => media.RowDisplay(
            media: v,
            leading: [Icon(mimex.icon(v.mimetype))],
            trailing: [
              media.ButtonShare(current: v),
              media.ButtonPlay(current: v),
            ],
          ),
          empty: ds.FileDropWell(upload),
        ),
      ),
    );
  }
}
