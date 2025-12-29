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

Geoget ist ein Werkzeug, das eine einfache Möglichkeit bietet, die aktuelle Vorabversion von PC/GEOS (https://github.com/bluewaysw/pcgeos) in Kombination mit der Basebox-Version (https://github.com/bluewaysw/pcgeos-basebox) zu testen.

ACHTUNG: Daten, die Sie mit PC/GEOS Ensemble erstellen, bleiben bei einem Update nicht erhalten. Dies ist ausschließlich für Test- und Debugging-Zwecke gedacht.

### Nutzung

Der einfachste Weg ist, geoget einfach zu starten, z. B. geoget-linux oder geoget-win64.exe aus dem Explorer. Danach finden Sie in Ihrem Home-Verzeichnis einen Ordner namens "geospc". In dem neu angelegten Verzeichnis können Sie den Launcher "ensemble.cmd" starten.

Wenn Sie mehr wollen, sieht die erweiterte Nutzung so aus:

```
geoget [Optionen] [install_root]

Optionen:
  -f, --force            vorhandene Installation ohne Rückfrage überschreiben
  -g, --geos <issue>     CI-latest-<issue> für GEOS-Downloads verwenden (akzeptiert 829 oder #829)
  -b, --basebox <issue>  CI-latest-<issue> für Basebox-Downloads verwenden (akzeptiert 13 oder #13)
  -h, --help             diese Hilfe anzeigen
  -l, --lang <lang>      nicht-englische GEOS-Sprache installieren (derzeit nur "gr" unterstützt)

Argumente:
  install_root           optionales Installationsverzeichnis; Standard ist "geospc" im Home-Verzeichnis

Standardverhalten:
  Wenn keine Issue-Optionen angegeben werden, wird CI-latest verwendet.
```
