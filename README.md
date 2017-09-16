[![Go Report Card](https://goreportcard.com/badge/github.com/31z4/harvest)](https://goreportcard.com/report/github.com/31z4/harvest)
[![Coverage Status](https://coveralls.io/repos/github/31z4/harvest/badge.svg?branch=master)](https://coveralls.io/github/31z4/harvest?branch=master)
[![CircleCI](https://circleci.com/gh/31z4/harvest.svg?style=svg)](https://circleci.com/gh/31z4/harvest)

# Harvest

`harvest` helps you understand what's inside your Redis. It samples Redis keys and shows top key prefixes.

The tool is designed to be simple and efficient. It does not require Redis [`DEBUG OBJECT`](https://redis.io/commands/debug-object) command to be available. Which is good if you're using AWS ElastiCache.

`harvest` should work with any version of Redis. Although, it's tested only with the latest stable version.

## Usage

Here is an example of how to use `harvest` and understand its output. Assuming Redis hostname is `redis` and it listens on a standard port `6379`.

```console
$ harvest redis://redis
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

The first line of the output warns us that the number of samples is greater than the database size. Which basically means that each key will be sampled multiple times. In future releases [`SCAN`](https://redis.io/commands/scan) will be used instead of [`RANDOMKEY`](https://redis.io/commands/randomkey) in that case.

The second line tells that from `1000` randomly sampled keys `648` start from `_kombu.binding.` and it's `33.98%` of all samples. The rest of the lines follow same pattern.

### Docker

You can use [31z4/harvest](https://hub.docker.com/r/31z4/harvest/) Docker image to run the tool. Assuming Redis container name is `redis` and it listens on a standard port `6379`.

```console
$ docker run --link redis:redis -it --rm 31z4/harvest redis://redis
```

## Feedback

Feedback is greatly appreciated. At this stage, the maintainers are most interested in feedback centered on the user experience (UX) of the tool. Do you have cases that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results?

## Contributing

Contributions are greatly appreciated. The maintainers actively manage the issues list, and try to highlight issues suitable for newcomers. The project follows the typical GitHub pull request model. See [CONTRIBUTING.md](CONTRIBUTING.md) for more details. Before starting any work, please either comment on an existing issue, or file a new one.

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details.
