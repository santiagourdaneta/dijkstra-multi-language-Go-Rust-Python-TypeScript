type Graph = { [key: string]: { [key: string]: number } };

function dijkstra(graph: Graph, start: string) {
    const distances: { [key: string]: number } = {};
    for (let node in graph) distances[node] = Infinity;
    distances[start] = 0;

    let queue = [{ node: start, dist: 0 }];

    while (queue.length > 0) {
        queue.sort((a, b) => a.dist - b.dist);
        const { node: currNode, dist: currDist } = queue.shift()!;

        if (currDist > distances[currNode]) continue;

        for (let neighbor in graph[currNode]) {
            let newDist = currDist + graph[currNode][neighbor];
            if (newDist < distances[neighbor]) {
                distances[neighbor] = newDist;
                queue.push({ node: neighbor, dist: newDist });
            }
        }
    }
    return distances;
}

const graph = { A: { B: 4, C: 2 }, B: { C: 5 }, C: { B: 1 } };
console.log(dijkstra(graph, 'A'));