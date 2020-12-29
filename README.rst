About
=====

A CLI based utility to operate Grafana, such as list/export/import/delete dashboards, create/delete/list annotations, etc.

Usage
-----

**Build**

::

  # Linux
  go build
  # Compile on Linux for Windows
  GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build


**Provide Grafana Information with CLI Options**

::

  ./grafana_cli help
  ./grafana_cli --server http://localhost:3000 --user admin --password admin dashboard list

**Provide Grafana Information with Env Vars**

::

  export GRAFANA_SERVER='http://localhost:3000'
  export GRAFANA_USER='admin'
  export GRAFANA_PASSWORD='admin'
  ./grafana_cli dashboard --help
  ./grafana_cli dashbaord export --help
  ./grafana_cli dashboard export --uid 'qpQbDyxMk' --path dashboard1.json
