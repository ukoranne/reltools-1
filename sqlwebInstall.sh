#!/bin/bash
cd /opt/flexswitch
tar -xvzf /opt/flexswitch/dbBrowser.tgz

# This is not needed... :-/
#cd /opt/flexswitch/dbbrowser
#dpkg -i /opt/flexswitch/dbbrowser/python-pkg-resources_20.1.1-1_all.deb

cd /opt/flexswitch/dbbrowser
tar -xzvf setuptools-20.1.1.tar.gz

cd /opt/flexswitch/dbbrowser/setuptools-20.1.1/
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf Pygments-2.1.3.tar.gz

cd /opt/flexswitch/dbbrowser/Pygments-2.1.3/
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf peewee-2.8.0.tar.gz

cd /opt/flexswitch/dbbrowser/peewee-2.8.0/
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf Flask-0.10.1.tar.gz

cd /opt/flexswitch/dbbrowser/Flask-0.10.1
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf itsdangerous-0.24.tar.gz

cd /opt/flexswitch/dbbrowser/itsdangerous-0.24/
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf Jinja2-2.4.tar.gz

cd /opt/flexswitch/dbbrowser/Jinja2-2.4
python setup.py install

cd /opt/flexswitch/dbbrowser
tar -xvzf Werkzeug-0.7.tar.gz

cd /opt/flexswitch/dbbrowser/Werkzeug-0.7
python setup.py install

cd /opt/flexswitch/dbbrowser/
tar -xvzf sqliteBrowser.tgz

cd /opt/flexswitch/dbbrowser/sqlite-web
python setup.py install
