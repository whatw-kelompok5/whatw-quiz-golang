# What Happen Around The World

How to run this project:

1. Run `go mod tidy` command
2. Run `go install github.com/cosmtrek/air@latest` (if want using air, If you don't want to use this, you can skip this step)
3. Run `go run main.go` or `air` for running using air


## step to acces endpoint :

### Quesions

#### GetAll Questions

**questions:**
`http://localhost:8083/api/questions`

response: 

```
{
    "data": [
        {
            "id": 2,
            "difficulty": "easy",
            "question": "Monas dibangun pada tahun berapa?",
            "options": [
                "17 Agustus 1961",
                "17 Agustus 1950",
                "17 Agustus 1945",
                "17 Agustus 1949"
            ],
            "answer": "17 Agustus 1961"
        },
        {
            "id": 3,
            "difficulty": "easy",
            "question": "Monas Berada di mana?",
            "options": [
                "Jakarta",
                "Bandung",
                "Surabaya",
                "Malang"
            ],
            "answer": "Jakarta"
        }
    ]
}

```
### Avatars

#### GetAll Avatars

**avatar:**
`http://localhost:8083/api/avatars`

response:

```
{
    "data": [
        {
            "id": 15,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1701080476/whatw-quiz/ond5gfhxe6huydwbapsp.jpg",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 150
            }
        },
        {
            "id": 13,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151875/whatw-quiz/ba3vdvz6ey9ygjd0bpgh.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        },
        {
            "id": 12,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151861/whatw-quiz/axybbjuikvl5ueuhhiag.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        },
        {
            "id": 11,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151800/whatw-quiz/twf4tjxzomrmsczjmswx.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        },
        {
            "id": 10,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151782/whatw-quiz/nxdxs3u9o5gprua9sxzj.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        },
        {
            "id": 9,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151767/whatw-quiz/bswzniegllk9ry5n06ae.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        },
        {
            "id": 8,
            "image": "https://res.cloudinary.com/dtha7yn1x/image/upload/v1700151750/whatw-quiz/rd1zvvu7t21ad7vitmoe.png",
            "price": {
                "@type": "type.googleapis.com/google.protobuf.Int32Value",
                "value": 0
            }
        }
    ]
}

```
