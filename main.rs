use std::collections::{BinaryHeap, HashMap};
use std::cmp::Ordering;

#[derive(Clone, Eq, PartialEq)]
struct State { cost: usize, position: String }

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.cost.cmp(&self.cost) // Min-heap
    }
}
impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> { Some(self.cmp(other)) }
}

fn main() {
    let mut graph = HashMap::new();
    graph.insert("A", vec![("B", 4), ("C", 2)]);
    graph.insert("B", vec![("C", 5)]);
    graph.insert("C", vec![("B", 1)]);

    let mut dists: HashMap<&str, usize> = graph.keys().map(|&k| (k, usize::MAX)).collect();
    let mut pq = BinaryHeap::new();

    dists.insert("A", 0);
    pq.push(State { cost: 0, position: "A".to_string() });

    while let Some(State { cost, position }) = pq.pop() {
        if cost > *dists.get(position.as_str()).unwrap_or(&usize::MAX) { continue; }
        if let Some(neighbors) = graph.get(position.as_str()) {
            for (neighbor, weight) in neighbors {
                let next_dist = cost + weight;
                if next_dist < *dists.get(neighbor).unwrap_or(&usize::MAX) {
                    dists.insert(neighbor, next_dist);
                    pq.push(State { cost: next_dist, position: neighbor.to_string() });
                }
            }
        }
    }
    println!("{:?}", dists);
}