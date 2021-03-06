package offline

// SystemTask create task
type SystemTask struct {
	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" bson:"_id" json:"identity"`
	Size          int64  `protobuf:"varint,2,opt,name=size,proto3" bson:"size" json:"size"`
	CreateUser    int64  `protobuf:"varint,3,opt,name=create_user,json=createUser,proto3" bson:"create_user" json:"createUser"`
	CreateTime    int64  `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" bson:"create_time" json:"createTime"`
	Name          string `protobuf:"bytes,5,opt,name=name,proto3" bson:"name" json:"name"`
	Type          int32  `protobuf:"varint,6,opt,name=type,proto3" bson:"type" json:"type"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" bson:"status" json:"status"`
	Flag          int32  `protobuf:"varint,8,opt,name=flag,proto3" bson:"flag" json:"flag"`
	ProcessedSize int64  `protobuf:"varint,9,opt,name=processed_size,json=processedSize,proto3" bson:"processed_size" json:"processedSize"`
	ErrorCode     int32  `protobuf:"varint,10,opt,name=error_code,json=errorCode,proto3" bson:"error_code" json:"errorCode"`
	ErrorMessage  string `protobuf:"bytes,11,opt,name=error_message,json=errorMessage,proto3" bson:"error_message" json:"errorMessage"`
	CreateAddr    string `protobuf:"bytes,12,opt,name=create_addr,json=createAddr,proto3" bson:"create_addr" json:"createAddr"`
	Data          string `protobuf:"bytes,13,opt,name=data,proto3" bson:"data" json:"data"`
	TextLink      string `protobuf:"bytes,14,opt,name=text_link,json=textLink,proto3" bson:"text_link" json:"textLink"`
	FileHash      string `protobuf:"bytes,18,opt,name=file_hash,json=fileHash,proto3" bson:"file_hash" json:"fileHash"`
	Username      string `protobuf:"bytes,19,opt,name=username,proto3" bson:"username" json:"username"`
	Password      string `protobuf:"bytes,20,opt,name=password,proto3" bson:"password" json:"password"`
}

// TaskFile the task stored in file
type TaskFile struct {
	Identity     string `protobuf:"bytes,1,opt,name=identity,json=identity,proto3" bson:"_id" json:"identity"`
	TaskIdentity string `protobuf:"bytes,2,opt,name=download_identity,json=downloadIdentity,proto3" bson:"task_identity" json:"downloadIdentity"`
	// PathIdentity string `protobuf:"bytes,2,opt,name=path_identity,json=pathIdentity,proto3" bson:"path_identity"`
	CreateTime    int64  `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" bson:"create_time" json:"create_time"`
	Name          string `protobuf:"bytes,5,opt,name=name,proto3" bson:"name" json:"name"`
	Path          string `protobuf:"bytes,6,opt,name=path,proto3" bson:"path" json:"path"`
	Hash          string `protobuf:"bytes,7,opt,name=hash,proto3" bson:"hash" json:"hash"`
	Size          int64  `protobuf:"varint,8,opt,name=size,proto3" bson:"size" json:"size"`
	ProcessedSize int64  `protobuf:"varint,9,opt,name=processed_size,json=processedSize,proto3" bson:"processed_size" json:"processedSize"`
	Status        int32  `protobuf:"varint,10,opt,name=status,proto3" bson:"status" json:"status"`
	Flag          int32  `protobuf:"varint,11,opt,name=flag,proto3" bson:"flag" json:"flag"`
	FileIndex     int32  `protobuf:"varint,12,opt,name=file_index,json=fileIndex,proto3" bson:"file_index" json:"fileIndex"`
	Finish        bool   `protobuf:"varint,13,opt,name=finish,proto3" bson:"finish" json:"finish"`
	Directory     bool   `protobuf:"varint,14,opt,name=directory,proto3" bson:"directory" json:"directory"`
}
