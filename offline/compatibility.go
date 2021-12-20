package offline

// SystemTask create task
type SystemTask struct {
	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" bson:"_id"`
	Size          int64  `protobuf:"varint,2,opt,name=size,proto3" bson:"size"`
	CreateUser    int64  `protobuf:"varint,3,opt,name=create_user,json=createUser,proto3" bson:"create_user"`
	CreateTime    int64  `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" bson:"create_time"`
	Name          string `protobuf:"bytes,5,opt,name=name,proto3" bson:"name"`
	Type          int32  `protobuf:"varint,6,opt,name=type,proto3" bson:"type"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" bson:"status"`
	Flag          int32  `protobuf:"varint,8,opt,name=flag,proto3" bson:"flag"`
	ProcessedSize int64  `protobuf:"varint,9,opt,name=processed_size,json=processedSize,proto3" bson:"processed_size"`
	ErrorCode     int32  `protobuf:"varint,10,opt,name=error_code,json=errorCode,proto3" bson:"error_code"`
	ErrorMessage  string `protobuf:"bytes,11,opt,name=error_message,json=errorMessage,proto3" bson:"error_message"`
	CreateAddr    string `protobuf:"bytes,12,opt,name=create_addr,json=createAddr,proto3" bson:"create_addr"`
	Data          string `protobuf:"bytes,13,opt,name=data,proto3" bson:"data"`
	TextLink      string `protobuf:"bytes,14,opt,name=text_link,json=textLink,proto3" bson:"text_link"`
	FileHash      string `protobuf:"bytes,18,opt,name=file_hash,json=fileHash,proto3" bson:"file_hash"`
	Username      string `protobuf:"bytes,19,opt,name=username,proto3" bson:"username"`
	Password      string `protobuf:"bytes,20,opt,name=password,proto3" bson:"password"`
}

// TaskFile the task stored in file
type TaskFile struct {
	Identity     string `protobuf:"bytes,1,opt,name=identity,json=identity,proto3" bson:"_id"`
	TaskIdentity string `protobuf:"bytes,2,opt,name=download_identity,json=downloadIdentity,proto3" bson:"task_identity"`
	// PathIdentity string `protobuf:"bytes,2,opt,name=path_identity,json=pathIdentity,proto3" bson:"path_identity"`
	CreateTime    int64  `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" bson:"create_time"`
	Name          string `protobuf:"bytes,5,opt,name=name,proto3" bson:"name"`
	Path          string `protobuf:"bytes,6,opt,name=path,proto3" bson:"path"`
	Hash          string `protobuf:"bytes,7,opt,name=hash,proto3" bson:"hash"`
	Size          int64  `protobuf:"varint,8,opt,name=size,proto3" bson:"size"`
	ProcessedSize int64  `protobuf:"varint,9,opt,name=processed_size,json=processedSize,proto3" bson:"processed_size"`
	Status        int32  `protobuf:"varint,10,opt,name=status,proto3" bson:"status"`
	Flag          int32  `protobuf:"varint,11,opt,name=flag,proto3" bson:"flag"`
	FileIndex     int32  `protobuf:"varint,12,opt,name=file_index,json=fileIndex,proto3" bson:"file_index"`
	Finish        bool   `protobuf:"varint,13,opt,name=finish,proto3" bson:"finish"`
	Directory     bool   `protobuf:"varint,14,opt,name=directory,proto3" bson:"directory"`
}
