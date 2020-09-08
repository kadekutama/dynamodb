# DynamoDB

This is a unit-testable version of [https://github.com/guregu/dynamo](https://github.com/guregu/dynamo). Basically, this version only wraps all functions to their respective interface.

## Example Code
```go
func InitDynamoDB() dynamodb.DB {
	return dynamodb.New(session.New(), aws.NewConfig())
}

func AnyLogic(db dynamodb.DB) (int64, error) {
	return db.Table("user").Get("user_id", "1").Count()
}
```

## Example Unit Test
```go
func TestAnyLogic(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockDB := mock.NewMockDB(mockCtrl)
	mockDB.EXPECT().Table("user").Return(
		func() dynamodb.Table {
			mockTable := mock.NewMockTable(mockCtrl)
			mockTable.EXPECT().Get("user_id", "1").Return(
				func() dynamodb.Query {
					mockQuery := mock.NewMockQuery(mockCtrl)
					mockQuery.EXPECT().Count().Return(int64(1), nil)

					return mockQuery
				}(),
			)

			return mockTable
		}(),
	)

	count, err := AnyLogic(mockDB)
	if err != nil || count != 1 {
		t.FailNow()
	}
}
```
