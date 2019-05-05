package dsa

// 顶点
type Vertex struct {
	Num  int
	Info interface{}
}

// 邻接矩阵(Adjacency Matrix)表示法
type MGraph struct {
	N        int      // 顶点数
	E        int      // 边数
	Edges    [][]int  // 所有边
	Vertices []Vertex // 所有结点
}

type ArcNode struct {
	ArcNode     Vertex      // 边表结点
	NextArcNode *ArcNode    // 边表下一个结点
	Info        interface{} // 边信息，例如权重
}

type VNode struct {
	VNode        Vertex   // 顶点表结点
	FirstArcNode *ArcNode // 边表第一个结点
}

// 邻接表(Adjacency List)表示法
type AGraph struct {
	N      int      // 顶点数
	E      int      // 边数
	VNodes []*VNode // 顶点表
}

func NewAGraph() *AGraph {
	return &AGraph{}
}

func (g *AGraph) DFS() {

}

// 判空
func (g *AGraph) IsEmpty() bool {
	return len(g.VNodes) == 0
}

// 增加一条边
func (g *AGraph) AddEdge(v1, v2 *Vertex) {

}

// TODO: Prim Algorithm
// TODO: Kruskal Algorithm
// TODO: Dijkstra Algorithm
// TODO: Floyd Algorithm
