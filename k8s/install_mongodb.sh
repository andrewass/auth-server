#!/bin/bash

helm install postgresql \
  --set postgresqlDatabase=stockdb \
  bitnami/postgresql