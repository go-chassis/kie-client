# kie client
the rest client of https://github.com/apache/servicecomb-kie

# Usage
new client
```go
	c, _ := kie.NewClient(kie.Config{
		Endpoint: "http://127.0.0.1:30110",
	})

```
create key value
```go
	kv := kie.KVRequest{
		Key:       "app.properties",
		Status:    "enabled",
		Labels:    map[string]string{"service": "client"},
		Value:     "timeout: 1s",
		ValueType: "text",
	}
	result, err := c.Create(context.TODO(), kv, kie.WithProject("client_test"))
```
update 
```go
	result, err := c.Create(context.TODO(), kv, kie.WithProject("client_test"))
	assert.NoError(t, err)
	kv := kie.KVRequest{
		Status:    "enabled",
		Value:     "timeout: 2s",
	}
	kv.ID = result.ID
	_, err = c.Put(context.TODO(), kv, kie.WithProject("client_test"))
```
long polling key values
```go
		go func() {
			kvs, _, err = c.List(context.TODO(),
				kie.WithLabels(map[string]string{"service": "client"}),
				kie.WithGetProject("client_test"),
				kie.WithWait("10s"))
			assert.NoError(t, err)
			assert.Equal(t, "timeout: 2s", kvs.Data[0].Value)
		}()
		kv := kie.KVRequest{
			ID:     result.ID,
			Key:    "app.properties",
			Labels: map[string]string{"service": "client"},
			Value:  "timeout: 2s",
		}
		_, err := c.Put(context.TODO(), kv, kie.WithProject("client_test"))
		assert.NoError(t, err)
```
delete
```go
err = c.Delete(context.TODO(), kv.ID, kie.WithProject("client_test"))
```