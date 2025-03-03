import 'dart:io';
import 'package:flutter/material.dart';
import 'package:fractal/downloads.dart' as downloads;
// import 'package:fractal/discovery.dart' as discovery;
import 'package:fractal/settings.dart' as settings;
import 'package:fractal/designkit.dart' as ds;
import 'package:fractal/design.kit/theme.defaults.dart' as theming;
import 'package:fractal/mdns.dart' as mdns;

class MyHttpOverrides extends HttpOverrides {
  @override
  HttpClient createHttpClient(SecurityContext? context) {
    return super.createHttpClient(context)
      ..badCertificateCallback = (X509Certificate cert, String host, int port) {
        return host == "localhost" || host == Platform.localHostname;
      };
  }
}

void main() {
  HttpOverrides.global = MyHttpOverrides();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      darkTheme: ThemeData(
        brightness: Brightness.dark,
        cardTheme: CardTheme(margin: EdgeInsets.all(10.0)),
        extensions: [theming.Defaults.defaults],
      ),
      themeMode: ThemeMode.dark,
      home: ds.Full(
        child: DefaultTabController(
          length: 3,
          child: Scaffold(
            appBar: TabBar(
              tabs: [
                // Tab(icon: Icon(Icons.share)),
                Tab(icon: Icon(Icons.movie)),
                Tab(icon: Icon(Icons.download)),
                Tab(icon: Icon(Icons.settings)),
              ],
            ),
            body: mdns.MDNSDiscovery(
              TabBarView(
                children: [
                  // discovery.Display(),
                  ds.ErrorBoundary(Icon(Icons.movie)),
                  ds.ErrorBoundary(downloads.Display()),
                  ds.ErrorBoundary(settings.Display()),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
