**Geoget** is a tool which gives you an easy way to test pre-release versions of PC/GEOS (https://github.com/bluewaysw/pcgeos) in combination with our specialized DosBox-Staging fork called Basebox, optimized for use with PC/GEOS (https://github.com/bluewaysw/pcgeos-basebox). There are versions of **geoget** for Win 64, Win 32, Linux, Mac and RPi.

CAUTION: data you create with these downloaded PC/GEOS Ensemble builds won't be preserved automatically on updates. This tool is primarily targeted at testing / debbugging.

The easiest way to use **Geoget** is to just launch it, e.g. ```geoget-linux``` or ```geoget-win64.exe``` from your file manager of choice. After doing so you will find a folder called ```geospc``` in your home folder. You can start the launcher ```ensemble.cmd``` or ```ensemble.sh``` in this directory.

If you want more, advanced usage looks like this:

### How to use

``` geoget [options] [install_root] ```

#### Options:

  ``` -f, --force ```: Overwrite existing installation without prompt

  ``` -g, --geos <issue> ```: Use ```CI-latest-<issue>``` for GEOS
  downloads (accepts issue numbers with or without "#")

  ```-b, --basebox <issue>```: Use ```CI-latest-<issue>``` for Basebox
  downloads

  ```-h, --help```: Show this help message.

  ```-l, --lang <lang>```: Non-english GEOS language to install.
  (Only "gr" supported for now and only with CI-latest.)

#### Arguments:

  ```install_root ```: Optional install root; defaults to "geospc" under home

#### Defaults:

  If no flags or arguments are provided, the english CI-latest version of GEOS
  is downloaded into "geospc", along with the regular Basebox (that does not provide host
  internet access yet).

### Usage Samples

```
# Download geos build for "issue #829" and basebox build for "issue #13"
# into the "geos-hostapi-2" folder.
# This will give you the test version for the new freeform DPI drivers and
# improved host internet access, especially on Linux.

geoget-win64 -g 829 -b 13 geos-hostapi-2

# Download the "split version", where we test options for an update-safe
# distribution of PC/GEOS with the "issue 2" build of basebox that offers
# the original host api interface, providing host internet access.

geoget-linux -g 973 -b 2 geos-split-install

# Donwload the regular CI-latest version of GEOS along with the Basebox version
# that gives you host internet access, overwriting an older version in "geospc".

geoget-win32 -b 2 -f

```

===================================================================

Deutsche Version

===================================================================

**Geoget** ist ein Werkzeug, das Ihnen eine einfache Möglichkeit gibt, Vorabversionen von PC/GEOS (https://github.com/bluewaysw/pcgeos) in Kombination mit unserem spezialisierten DosBox-Staging-Fork namens "Basebox" zu testen, die für die Verwendung mit PC/GEOS optimiert ist (https://github.com/bluewaysw/pcgeos-basebox). Es gibt Versionen von **geoget** für Win 64, Win 32, Linux, Mac und RPi.

VORSICHT: Daten, die Sie mit diesen heruntergeladenen PC/GEOS Ensemble-Builds erstellen, bleiben bei Updates nicht erhalten. Dieses Tool ist hauptsächlich zum Testen / Debuggen gedacht.

Der einfachste Weg, **Geoget** zu verwenden, besteht darin, es einfach zu starten, z. B. ```geoget-linux``` oder ```geoget-win64.exe``` aus dem Dateimanager Ihrer Wahl. Danach finden Sie einen Ordner namens ```geospc``` in Ihrem Home-Ordner. Sie können den Launcher ```ensemble.cmd``` oder ```ensemble.sh``` in diesem Verzeichnis starten.

Wenn Sie mehr wollen, sieht die erweiterte Nutzung so aus:

### Wie man es benutzt

``` geoget [Optionen] [install_root] ```

#### Optionen:

  ``` -f, --force ```: Bestehende Installation ohne Aufforderung überschreiben

  ``` -g, --geos <issue> ```: Verwenden Sie ```CI-latest-<issue>``` für GEOS-Downloads (akzeptiert Problemnummern mit oder ohne "#")

  ```-b, --basebox <issue>```: Verwenden Sie ```CI-latest-<issue>``` für Basebox-Downloads

  ```-h, --help```: Diese Hilfemeldung anzeigen.

  ```-l, --lang <lang>```: Nicht-englische GEOS-Sprache zum Installieren. (Derzeit wird nur "gr" unterstützt und nur mit CI-latest.)

#### Argumente:

  ```install_root ```: Optionales Installationsverzeichnis; Standard ist "geospc" unter home

#### Standardeinstellungen:

  Wenn keine Flags oder Argumente angegeben werden, wird die englische CI-latest-Version von GEOS zusammen mit der regulären Basebox (die noch keinen Host-Internetzugang bietet) in "geospc" heruntergeladen.

### Anwendungsbeispiele

```
# Lade den Geos-Build für "Issue #829" und den Basebox-Build für "Issue #13"
# in den Ordner "geos-hostapi-2" herunter.
# Dies gibt Ihnen die Testversion für die neuen Freeform-DPI-Treiber und
# verbesserten Host-Internetzugang, insbesondere unter Linux.

geoget-win64 -g 829 -b 13 geos-hostapi-2

# Lade die "geteilte Version" herunter, in der wir Optionen für eine update-sichere
# Verteilung von PC/GEOS mit dem "Issue 2"-Build von Basebox testen, der
# die ursprüngliche Host-API-Schnittstelle bietet und Host-Internetzugang ermöglicht.

geoget-linux -g 973 -b 2 geos-split-install

# Lade die reguläre CI-latest-Version von GEOS zusammen mit der Basebox-Version herunter,
# die Ihnen Host-Internetzugang ermöglicht, und überschreibe eine ältere Version in "geospc".

geoget-win32 -b 2 -f

```
