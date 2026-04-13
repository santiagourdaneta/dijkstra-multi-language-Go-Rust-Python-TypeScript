# Algoritmo de Dijkstra Multi-Lenguaje 🧭

Este repositorio contiene la implementación del **Algoritmo de Dijkstra**, una pieza fundamental de las matemáticas aplicadas a la informática para encontrar el camino más corto entre nodos en un grafo ponderado.

## 🚀 Lenguajes Incluidos
* **Python**: Implementación legible usando `heapq`.
* **Go**: Enfoque de alto rendimiento con interfaces de `container/heap`.
* **Rust**: Gestión segura de memoria y uso de `BinaryHeap`.
* **TypeScript**: Lógica moderna para entornos web/Node.js.

## 🧠 ¿Cómo funciona?
El algoritmo utiliza una **Cola de Prioridad** para explorar siempre el nodo con la distancia acumulada más baja desde el origen, actualizando los costos de los vecinos hasta encontrar la ruta óptima.

## 🛠️ Ejecución
- **Python**: `python main.py`
- **Go**: `go run main.go`
- **TypeScript**: `node main.js`
- **Rust**: `rustc main.rs && ./main` 

## 📐 Matemáticas Detrás del Código: Complejidad
El Algoritmo de Dijkstra es un ejemplo clásico de optimización. Su eficiencia depende de cómo manejamos la frontera de exploración:

- **Complejidad Temporal:** $O((V + E) \log V)$ utilizando una **Cola de Prioridad** (Binary Heap), donde $V$ es el número de vértices y $E$ el de aristas.
- **¿Por qué usar Heaps?** Sin una cola de prioridad, tendríamos que buscar en toda la lista el nodo más cercano en cada paso ($O(V^2)$). El "Heap" reduce esa búsqueda de lineal a logarítmica, permitiendo que aplicaciones como Google Maps procesen millones de intersecciones en segundos.