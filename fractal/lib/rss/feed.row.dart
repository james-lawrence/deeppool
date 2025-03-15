import 'package:flutter/material.dart';
import 'package:fractal/designkit.dart' as ds;
import './rss.pb.dart';

class FeedRow extends StatelessWidget {
  final Feed current;
  final Function(Feed)? onChange;
  FeedRow({super.key, Feed? current, this.onChange})
    : current = current ?? (Feed.create()..autodownload = false);

  @override
  Widget build(BuildContext context) {
    final themex = ds.Defaults.of(context);

    return Row(
      spacing: themex.spacing ?? 0.0,
      children: [
        if (current.hasDescription()) SelectableText(current.description),
        SelectableText(current.url),
        Spacer(),
        current.autodownload ? Icon(Icons.downloading_rounded) : Container(),
        current.autoarchive ? Icon(Icons.archive_outlined) : Container(),
      ],
    );
  }
}
