package typ

type Sheet struct {
	Name string `json:"name"`
	Rows []Row  `json:"rows"`
}

type Row struct {
	Axis  string      `json:"axis"`
	Value interface{} `json:"value"`
}

type ChunkData struct {
	TaskId    string `json:"task_id"`
	SheetName string `json:"sheet_name"`
	Rows      []Row  `json:"rows"`
}
