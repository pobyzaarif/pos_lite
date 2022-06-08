### Filebeat Setup / Configuration

- edit `filebeat.yml`
```
filebeat.inputs:
- type: log
  enabled: true
  paths:
    - /Users/7noob/Project/pos_lite/logs/*
  json.keys_under_root: false
  json.add_error_key: true

setup.template.name: "filebeat"
setup.template.fields: "fields-custom.yml"
setup.template.overwrite: true

processors:
  - decode_json_fields:
      fields: [message"]
      process_array: false
      max_depth: 1
      target: "ya"
      overwrite_keys: false
      add_error_key: true
```

- add new file `fields-custom.yml`
```
- key: my-custom-beat
  title: my-custom-beat
  description: These are the fields used by my-custom-beat.
  fields:
    - name: '@timestamp'
      level: core
      required: true
      type: date
      description: 'Date/time when the event originated.'
      example: '2016-05-23T08:05:34.853Z'
    - name: json.caller
      type: keyword
      required: true
    - name: json.error
      type: keyword
      required: true
    - name: json.level
      type: keyword
      required: true
    - name: json.message
      type: keyword
      required: true
    - name: json.processing_time
      type: integer
      required: true
    - name: json.service_name
      type: keyword
      required: true
    - name: json.tag
      type: keyword
      required: true
    - name: json.time
      type: keyword
      required: true
    - name: json.timer_end
      type: keyword
      required: true
    - name: json.timer_start
      type: keyword
      required: true
    - name: json.tracker_id
      type: keyword
      required: true
    - name: json.data
      type: object
    - name: json.data.panic.trace
      type: text
    - name: json.data.net.handler
      type: keyword
    - name: json.data.net.host
      type: keyword
    - name: json.data.net.method
      type: keyword
    - name: json.data.net.remote_ip
      type: keyword
    - name: json.data.net.request
      type: text
    - name: json.data.net.request_header
      type: text
    - name: json.data.net.response
      type: text
    - name: json.data.net.response_header
      type: text
    - name: json.data.net.response_http_code
      type: short
    - name: json.data.net.url
      type: keyword
    - name: json.data.db.query
      type: text
    - name: json.data.db.rows
      type: integer
    - name: json.data.app
      type: object
      enabled: false
```