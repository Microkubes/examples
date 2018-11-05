Micro Todo
==========
**Version:** 

### /todo/{todoID}
---
##### ***GET***
**Summary:** getById todo

**Description:** Get todo by ID

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| todoID | path | Todo ID | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [TodoMedia](#todomedia) |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

##### ***PATCH***
**Summary:** updateTodo todo

**Description:** Update todo

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| todoID | path | Todo ID | Yes | string |
| payload | body | Todo update payload | Yes | [TodoUpdatePayload](#todoupdatepayload) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [TodoMedia](#todomedia) |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

##### ***PUT***
**Summary:** updateTodo todo

**Description:** Update todo

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| todoID | path | Todo ID | Yes | string |
| payload | body | Todo update payload | Yes | [TodoUpdatePayload](#todoupdatepayload) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [TodoMedia](#todomedia) |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

### /todo/{todoID}/delete
---
##### ***DELETE***
**Summary:** deleteTodo todo

**Description:** Delete todo

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| todoID | path | Todo ID | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK |  |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

### /todo/add
---
##### ***POST***
**Summary:** addTodo todo

**Description:** Add new todo

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| payload | body | Todo payload | Yes | [TodoPayload](#todopayload) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [TodoMedia](#todomedia) |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

### /todo/all
---
##### ***GET***
**Summary:** getAllTodos todo

**Description:** Get all todos

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query | Limit todos per page | No | integer |
| offset | query | number of todos to skip | No | integer |
| order | query | order by | No | string |
| sorting | query |  | No | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK |  |
| 400 | Bad Request | [error](#error) |
| 404 | Not Found | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

### /todo/filter
---
##### ***POST***
**Summary:** filterTodos todo

**Description:** Filter (lookup) todos

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| payload | body |  | Yes | [FilterTodoPayload](#filtertodopayload) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [PaginatedTodosMedia](#paginatedtodosmedia) |
| 400 | Bad Request | [error](#error) |
| 500 | Internal Server Error | [error](#error) |

### Models
---

### FilterTodoPayload  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| filter |  | Filter by fields key=>value | No |
| order | [ [OrderSpecs](#orderspecs) ] | Sort specifications. | No |
| page | long | Page number to fetch | Yes |
| pageSize | long | Number of items per page | Yes |

### OrderSpecs  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| direction | string | Sort direction. One of 'asc' (ascending) or 'desc' (descenting). | No |
| property | string | Order by property | No |

### PaginatedTodosMedia  

PaginatedTodosMedia media type (default view)

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| items | [ [TodoMedia](#todomedia) ] | List of todos | No |
| page | long | Current page number | No |
| pageSize | long | Number of items per page | No |
| total | long | Total number of items | No |

### TodoMedia  

TodoMedia media type (default view)

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| completedAt | long |  | No |
| createdAt | long |  | Yes |
| createdBy | string |  | No |
| description | string |  | No |
| done | boolean |  | No |
| id | string |  | Yes |
| title | string |  | No |

### TodoPayload  

Todo payload

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string | Todo description | Yes |
| title | string | Todo title | Yes |

### TodoUpdatePayload  

Todo update payload

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string | Todo description | No |
| done | boolean | Todo status | No |
| title | string | Todo title | No |

### error  

Error response media type (default view)

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | string | an application-specific error code, expressed as a string value. | No |
| detail | string | a human-readable explanation specific to this occurrence of the problem. | No |
| id | string | a unique identifier for this particular occurrence of the problem. | No |
| meta | object | a meta object containing non-standard meta-information about the error. | No |
| status | string | the HTTP status code applicable to this problem, expressed as a string value. | No |