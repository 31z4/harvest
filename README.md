[![Go Report Card](https://goreportcard.com/badge/github.com/31z4/harvest)](https://goreportcard.com/report/github.com/31z4/harvest)
[![Coverage Status](https://coveralls.io/repos/github/31z4/harvest/badge.svg?branch=master)](https://coveralls.io/github/31z4/harvest?branch=master)
[![CircleCI](https://circleci.com/gh/31z4/harvest.svg?style=svg)](https://circleci.com/gh/31z4/harvest)

# Harvest

`harvest` helps you understand what's inside your Redis by sampling its keys and returning insightful stats about it.

```
warning: database size (12) is less than the number of samples (1000)

_kombu.binding.: 33.98% (648)
unacked: 9.49% (181)
_kombu.binding.celery: 9.18% (175)
_kombu.binding.reply.celery.pidbox: 6.45% (123)
test: 5.98% (114)
_kombu.binding.schedule: 5.35% (102)
_kombu.binding.blocks: 5.19% (99)
_kombu.binding.default: 4.82% (92)
_kombu.binding.celery.pidbox: 4.72% (90)
_kombu.binding.celeryev: 4.46% (85)
```
