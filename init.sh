#!/bin/bash

psql -d ${POSTGRES_DB} -U ${POSTGRES_USER} -f /tmp/dump.sql
