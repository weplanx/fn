# Func API

Provide high-performance function extension services

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kainonly/func-api?style=flat-square)](https://github.com/kainonly/func-api)
[![Github Actions](https://img.shields.io/github/workflow/status/kainonly/func-api/release?style=flat-square)](https://github.com/kainonly/func-api/actions)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/func-api?style=flat-square)](https://hub.docker.com/r/kainonly/func-api)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/func-api.svg?style=flat-square)](https://hub.docker.com/r/kainonly/func-api)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/func-api/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.8"
services: 
    func:
        image: kainonly/func-api
        restart: always
        volumes: 
            - ./func/config:/app/config
            - ./func/fonts:/app/fonts
        ports: 
            - 6000:6000
```

## Configuration

For configuration, please refer to `config/config.example.yml` and create `config/config.yml`

- **debug** `string` turn on debugging, that is `net/http/pprof`, and visit the address `http://localhost: 6060/debug/pprof`
- **listen** `string` listening address
- **logger** `bool` turn on logger
- **database** object buffer database
  - **dns** `string` user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
  - **max_idle_conns** `int` sets the maximum number of connections in the idle connection pool
  - **max_open_conns** `int` sets the maximum number of open connections to the database
  - **conn_max_lifetime** `int` sets the maximum amount of time a connection may be reused
  - **table_prefix** `string` database table prefix
- **storage** storage settings
  - **drive** `string` support *local*(Local storage) *oss*(Alibaba Cloud Object Storage) *obs*(Huawei Cloud Object Storage) *cos*(Tencent Cloud Object Storage)
  - **option** refer to `config/config.example.yml`
- **fonts** `object` font library

## API

The `debug.FreeOSMemory()` is used in the service, maybe this is not recommended by the application, but a large number of file processing does need to release the memory in time

### POST /excel/simple

Excel is easy to generate, scene description: small amount of data, customizable (in the future, table mergers, functions, table styles, charts, etc. will be added)

#### parameters (JSON):

- **sheets** `any[]` 
  - **name** `string` worksheet name
  - **rows** `any[]` set cells value 
    - **axis** `string` E.g. A1
    - **value** `any`

#### response body (JSON):

- **url** `string` generate save path

### POST /excel/new_task

Excel chunk task, scene description: suitable for mass data generation

#### parameters (JSON):

- **sheets_name** `string[]` worksheets name

#### response body (JSON):

- **task_id** `string` task id

### POST /excel/add_row

Add rows to the task

#### parameters (JSON):

- **task_id** `string` task id
- **sheet_name** `string` worksheet name
- **rows** `any[]` set cells value 
  - **axis** `string` E.g. A1
  - **value** `any`

### POST /excel/commit_task

Submit task

#### parameters (JSON):

- **task_id** `string` task id

#### response body (JSON):

- **url** `string` generate save path

### POST /qrcode/testing

QR code generation

#### parameters (JSON):

- **content** `string` QR code content
- **size** `number` QR code size
- **text** `any[]` Text settings
  - **value** `string` text value
  - **type** `string` text type
  - **size** `number` text size
  - **margin** `number` margin

#### response body:

png image of base64 content

### POST /qrcode/pre_built

QR code pre-built

#### parameters (JSON):

- **options** `any[]`
  - **content** `string` QR code content
  - **size** `number` QR code size
  - **text** `any[]` Text settings
    - **value** `string` text value
    - **type** `string` text type
    - **size** `number` text size
    - **margin** `number` margin
- **update** `bool` mandatory update

### POST /qrcode/export

Export to excel

#### parameters (JSON):

- **lists** `string[]` Multiple QR code content
- **height** `float` Row height

#### response body (JSON):

- **url** `string` generate save path