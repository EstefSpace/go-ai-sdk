# Go-Ai-Sdk

## 🇪🇸 Español 
<hr>
La libreria más completa con varios modelos de inteligencia artificial, facil de instalar y de usar. Para todos los modelos se **requiere de una API KEY**.

## Documentación
<hr>
La libreria cuenta con una documentación detallada de como es que se pueda usar los modelos que se ofrece. Los errores que pueden llegar a ocurrir, etc. Visitalo haciendo click <a href="https://goaisdk.info">aqui</a>.

## Ejemplo de Uso
Aqui se puede ver un pequeño ejemplo de uso utilizando un modelo de Gemini:

```go
    package main

import (
	"fmt"
	"github.com/estefspace/go-ai-sdk"
)

func main() {
	prompt := "Cuentame la historia de Chicharito"
    instructions := "Te llamas AI History y eres experto en temas de historia, de personajes de todo de historia."
	client := gemini.NewClient("API_KEY", "MODELO") //ver documentación

	content, err := client.Ask(prompt, instructions)
	if err != nil {
		fmt.Println("Ocurrio un error")
        fmt.Println(err)
	}

	fmt.Println(content)
}
```

<hr>

##  🇺🇸 English
<hr>
The most comprehensive library with several artificial intelligence models, easy to install and use. All models **require an API key**.


## Documentation
<hr>
The library has detailed documentation on how to use the models it offers, including possible errors, and more. Visit it by clicking <a href="https://goaisdk.info">here</a>.

## Example of Use
Here you can see a small example of use using a Gemini model:

```go
    package main

import (
	"fmt"
	"github.com/estefspace/go-ai-sdk"
)

func main() {
	prompt := "How are you?"
    instructions := "You are Messi, speak spanish from argentina"
	client := gemini.NewClient("API_KEY", "MODEL") //view docs

	content, err := client.Ask(prompt, instructions)
	if err != nil {
		fmt.Println("Uppss error!")
        fmt.Println(err)
	}

	fmt.Println(content)
}
```
