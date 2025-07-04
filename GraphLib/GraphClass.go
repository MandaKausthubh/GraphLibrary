package graphlib


/*
	GraphLibrary for GeoGraphical Production and Supply Chain Management
*/

type GeoNode interface {
	getID() 		string
	getDBKey()		string
}

type EdgeWeight interface {
	getCost(key int) float32
}

type GeoEdge interface {
	getNodeFrom() 	string 
	getNodeTo() 	string 
	getEdgeWeight() EdgeWeight
}

type GeoGraph interface {
	AddNode(GeoNode)
	AddEdge(GeoEdge)
	
	AllPairsShortestPath(sourceNode int, Distances *[]int)
	MaxCut(sourceNodes *[]int, destinationNodes *[]int)
}

// Get an Create a super Regional Node
type GeoNodeRegional struct{
	NodeID string
	DB_info string
	Representation string
}

type GeoEgdeRegional struct {
	NodeID1, NodeID2 string
	ListOfPaths 	[]string
	CostOfPaths 	[]float32
	TimeOfPaths 	[]float32
}

type Pair struct {
	NodeIdFrom, NodeIDTo string
}

type GeoGraphRegional struct {
	GraphID 		string
	NumNodes 		int
	NumEdges 		int
	NodesList 		[]GeoNodeRegional
	EdgeList 		map[Pair]GeoEgdeRegional
}
