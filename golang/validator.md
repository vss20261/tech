# Go Validator
* [github](https://github.com/go-playground/validator)
* [reference](https://pkg.go.dev/github.com/go-playground/validator/v10)

---

* domain > validate
    - **Create**
        - 구조체를 따로 만들 필요없이 model의 태그를 정의해 validate
        - model의 정의된 태그는 create시 사용 되므로 create상황에 맞게 태그를 지정
          ~~~ go
            /* 
               1. create시 ID는 비어있기에 필수 태그 지정 x
               2. gorm: default 태그가 존재 한다면 required 태그 제거가능
               3. nullable 필드의 경우 omitempty 태그 사용
            */
            ex) type Role struct {
                  ID            uint
                  ClientID      uint      `validate:"required"`
                  Name          string    `validate:"required,lte=500"`
                  Explanation   *string   `validate:"omitempty,lte=500"`
                  Note          *string   `validate:"omitempty,lte=255"`
                  IsActive      *bool     `gorm:"default:false"`
                  CreatorID     uint      `validate:"required"`
                  DateCreated   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
                  LastUpdatorID uint      `validate:"required"`
                  LastUpdated   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
                }
          ~~~
    - **Update**
        - 기능별로 구조체 생성 후 validate
        - 구조체 naming - Partial{테이블 명}{기능}
        - ex ) type PartialFileWorkUpload struct {}
