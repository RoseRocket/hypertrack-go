# hypertrack-go

The Golang library for interacting with the Hypertrack.io HTTP API.

## Installation

```
$ go get github.com/roserocket/hypertrack-go
```

## Getting Started

```go
package main

import hypertrack "github.com/roserocket/hypertrack-go"

func main(){

	// instantiate a client
	client := hypertrack.NewClient("your_app_key", "your_app_secret")

    // create user
	user, err := client.CreateUser("John", "+16267777777", "http://your-photo-url", "user-lookup-id", "group-id")
	if err != nil {
		return err, nil
	}

}
```

## Methods

### Create User

#### `func (c *Client) CreateUser(name, phone, photo, lookupId, groupId string) (*User, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|name   | `string`  | The name of user   |
|phone  | `string`  | The phone number |
|photo  | `string`  | The photo |
|lookupId  | `string`  | The lookup id |
|groupId  | `string`  | The group id |

```go
user, err := client.CreateUser("John", "+16267777777", "http://your-photo-url", "user-lookup-id", "group-id")
if err != nil {
    return err, nil
}
```

### Retrieve User

#### `func (c *Client) RetrieveUser(userId string) (*User, error) {`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|userId   | `string`  | User Id   |

```go
user, err := client.RetrieveUser("user-id")
if err != nil {
    return err, nil
}
```

### Create Action (using address)

#### `func (c *Client) CreateActionUsingAddress(address, city, zipCode, country, lookupId string, actionType ActionType, scheduledAt *time.Time) (*Action, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|address   | `string`  | Address   |
|city  | `string`  | City |
|zipCode  | `string`  | Zip Code |
|country  | `string`  | Country |
|lookupId  | `string`  | The lookup id |
|actionType  | `string`  | Action type |
|scheduledAt  | `timestamp`  | ISO8601 timestamp |

```go
action, err := client.CreateAction("123 Fake St", "Toronto", "M5V 1C8", "Canada", "lookup-id", "pickup", "2017-06-02T18:14:04.481983Z")
if err != nil {
    return err, nil
}
```

### Complete Action

#### `func (c *Client) CompleteAction(actionId string) (*Action, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|actionId   | `string`  | Action Id   |

```go
action, err := client.CompleteAction("action-id")
if err != nil {
    return err, nil
}
```

### Cancel Action

#### `func (c *Client) CancelAction(actionId string) (*Action, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|actionId   | `string`  | Action Id   |

```go
action, err := client.CancelAction("action-id")
if err != nil {
    return err, nil
}
```

### Retrieve User

#### `func (c *Client) RetrieveAction(actionId string) (*User, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|actionId   | `string`  | Action Id   |

```go
action, err := client.RetrieveAction("action-id")
if err != nil {
    return err, nil
}
```

### Assign Actions to User

#### `func (c *Client) AssignActionToUser(userId string, actionIds []string) (*User, error)`

|Argument   | Type  | Description   |
|:-:|:-:|:-:|
|userId   | `string`  | User Id   |
|actionIds   | `[]string`  | Array of Action Ids   |

```go
user, err := client.AssignActionToUser("user-id", []string{"action-id1", "action-id2"})
if err != nil {
    return err, nil
}
```

## License

This code is free to use under the terms of the MIT license.