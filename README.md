# Kit
Set of common packages used by all or most projects in my repositories.

## Table of Contents

- [Implementation examples](#implementation-examples)
    - [Postgresql](#postgresql)
    - [Redis](#redis)
- [Testing](#testing)

## Implementation examples

### Postgresql 

See the examples already implemented in the folder `/postgresql`.

### Redis

See the examples already implemented in the folder `/redis`.

## Testing

To run the tests simply execute the following command:

```shell
make test
```

This will stop any containers defined by the compose file for tests if already running
and then rebuild the containers using the compose file.

To down the containers simply execute the following command:

```shell
make test-down
```

lairon14@gmail.com  
-- Lairon