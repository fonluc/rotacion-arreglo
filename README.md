# Rotación de Arreglo en Go

## Descripción

Este proyecto implementa una función en Go para rotar los elementos de un arreglo a la derecha un número dado de veces. La implementación se prueba con varios casos de ejemplo para verificar su funcionamiento.

## Implementación

La función `rotateRight` rota los elementos de un arreglo a la derecha `m` veces. La función maneja arreglos de cualquier tamaño, incluyendo arreglos vacíos, y realiza las rotaciones de manera eficiente.

### Función `rotateRight`

```go
func rotateRight(arr []int, m int) []int {
    n := len(arr)
    if n == 0 {
        return arr
    }
    m = m % n
    return append(arr[n-m:], arr[:n-m]...)
}
```

- **Parámetros:**
  - `arr`: El arreglo que se debe rotar.
  - `m`: El número de rotaciones a realizar.
- **Resultado:**
  - Un nuevo arreglo rotado a la derecha `m` veces.

### Función `printArray`

```go
func printArray(arr []int) {
    fmt.Print("arr[")
    for i, v := range arr {
        if i > 0 {
            fmt.Print(",")
        }
        fmt.Print(v)
    }
    fmt.Println("]")
}
```

- **Propósito:**
  - Imprimir el arreglo en un formato legible.

## Pruebas

El código incluye pruebas para verificar que la función `rotateRight` funcione correctamente con diferentes entradas. Se presentan los casos de prueba a continuación:

### Ejemplos de Entrada y Salida

| Ejemplo       | Valor de m | Resultado       |
|---------------|------------|-----------------|
| arr[1,2,3,4]  | m = 1      | arr[4,1,2,3]    |
| arr[1,2,3,4]  | m = 2      | arr[3,4,1,2]    |
| arr[1,2,3,4]  | m = 3      | arr[2,3,4,1]    |
| arr[1,2,3,4]  | m = 4      | arr[1,2,3,4]    |

### Explicación de las Pruebas

1. **Entrada: `arr[1,2,3,4]`, m = 1**
   - **Descripción:** Rota el arreglo una vez a la derecha.
   - **Resultado esperado:** `arr[4,1,2,3]`.

2. **Entrada: `arr[1,2,3,4]`, m = 2**
   - **Descripción:** Rota el arreglo dos veces a la derecha.
   - **Resultado esperado:** `arr[3,4,1,2]`.

3. **Entrada: `arr[1,2,3,4]`, m = 3**
   - **Descripción:** Rota el arreglo tres veces a la derecha.
   - **Resultado esperado:** `arr[2,3,4,1]`.

4. **Entrada: `arr[1,2,3,4]`, m = 4**
   - **Descripción:** Rota el arreglo cuatro veces a la derecha (equivalente a no rotar).
   - **Resultado esperado:** `arr[1,2,3,4]`.

## Cómo Funciona

1. **Definición de Casos de Prueba:**
   - Los casos de prueba están definidos en una lista con diferentes arreglos y valores de `m`.

2. **Ejecución de Rotaciones:**
   - Para cada caso de prueba, la función `rotateRight` rota el arreglo y se imprime el resultado.

3. **Formato de Salida:**
   - Los resultados se presentan en un formato tabulado que facilita la visualización de los arreglos rotados.
