[![Go Report Card](https://goreportcard.com/badge/github.com/31z4/harvest)](https://goreportcard.com/report/github.com/31z4/harvest)
[![Coverage Status](https://coveralls.io/repos/github/31z4/harvest/badge.svg?branch=master)](https://coveralls.io/github/31z4/harvest?branch=master)
[![CircleCI](https://circleci.com/gh/31z4/harvest.svg?style=svg)](https://circleci.com/gh/31z4/harvest)

# Harvest

`harvest` helps you understand what's inside your Redis. It samples Redis keys and returns some insightful stats. It particular it aims to answer two simple questions that are hard to answer using existing tools. The questions are:

1. What are top 10 key prefixes?
2. What are top 10 most memory consuming keys?

Answers to these questions should help you figure out why your Redis has suddenly started to eat so much memory or even reached its capacity.

The tool is designed to be simple and efficient. It does not require Redis [`DEBUG OBJECT`](https://redis.io/commands/debug-object) command to be available. Which is good if you're using AWS ElastiCache.

## Usage

Here is an example of how to use `harvest` with Docker and understand its output. Assuming Redis container name is `redis` and it listens on a standard port `6379`.

```console
$ docker run --link redis:redis -it --rm harvest redis://redis
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

The first line of the output warns us that the number of samples is greater than the database size. Which basically means that results may be more inaccurate.

The second line tells that from `1000` randomly sampled keys `648` start from `_kombu.binding.` and it's `33.98%` of all samples.

## Feedback

Feedback is greatly appreciated. At this stage, the maintainers are most interested in feedback centered on the user experience (UX) of the tool. Do you have cases that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results?

## Contributing

Contributions are greatly appreciated. The maintainers actively manage the issues list, and try to highlight issues suitable for newcomers. The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue, or file a new one.

## License

This project is licensed under the GPLv3 License - see the [LICENSE.md](LICENSE.md) file for details.
