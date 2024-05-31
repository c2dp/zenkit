package info

type ClusterSettings struct {
	ID          string      `yaml:"id"`
	Name        interface{} `yaml:"name"`
	Description interface{} `yaml:"description"`
	Master      string      `yaml:"master"`
	Hosts       []struct {
		IP       string `yaml:"ip"`
		Hostname string `yaml:"hostname"`
		Tag      int    `yaml:"tag"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Status   string `yaml:"status"`
		Health   string `yaml:"health"`
		Vars     struct {
			Layout struct {
				LogDir       string   `yaml:"log_dir"`
				SsdDirs      []string `yaml:"ssd_dirs"`
				RuntimeDir   string   `yaml:"runtime_dir"`
				ConfDir      string   `yaml:"conf_dir"`
				HomeDir      string   `yaml:"home_dir"`
				SataDirs     []string `yaml:"sata_dirs"`
				ZookeeperDir []string `yaml:"zookeeper_dir"`
				DataDirs     []string `yaml:"data_dirs"`
			} `yaml:"layout"`
			System struct {
				MemtotalMb          string `yaml:"memtotal_mb"`
				DistributionVersion string `yaml:"distribution_version"`
				DefaultIpv4         struct {
					Broadcast  string `yaml:"broadcast"`
					Address    string `yaml:"address"`
					Netmask    string `yaml:"netmask"`
					Alias      string `yaml:"alias"`
					Macaddress string `yaml:"macaddress"`
					Interface  string `yaml:"interface"`
					Type       string `yaml:"type"`
					Gateway    string `yaml:"gateway"`
					Mtu        int    `yaml:"mtu"`
					Network    string `yaml:"network"`
				} `yaml:"default_ipv4"`
				ProcessorVcpus           string `yaml:"processor_vcpus"`
				SwaptotalMb              string `yaml:"swaptotal_mb"`
				Distribution             string `yaml:"distribution"`
				DistributionMajorVersion string `yaml:"distribution_major_version"`
				ProductName              string `yaml:"product_name"`
				Architecture             string `yaml:"architecture"`
			} `yaml:"system"`
		} `yaml:"vars"`
		HostFact struct {
			CPUInfo struct {
				CPUUsage string `yaml:"cpuUsage"`
				CPUCores string `yaml:"cpuCores"`
			} `yaml:"cpuInfo"`
			MemoryInfo struct {
				TotalMemory     string `yaml:"total_memory"`
				UsedMemory      string `yaml:"used_memory"`
				MemoryUsageRate string `yaml:"memory_usage_rate"`
			} `yaml:"memoryInfo"`
			DiskInfo struct {
				Total struct {
					TotalUsed      string `yaml:"totalUsed"`
					TotalCap       string `yaml:"totalCap"`
					TotalUsageRate string `yaml:"totalUsageRate"`
				} `yaml:"total"`
				DiskDetail []struct {
					Disk     string `yaml:"disk"`
					Usage    string `yaml:"usage"`
					Used     string `yaml:"used"`
					Dir      string `yaml:"dir"`
					Message  string `yaml:"message"`
					Capacity string `yaml:"capacity"`
					Status   string `yaml:"status"`
				} `yaml:"diskDetail"`
			} `yaml:"diskInfo"`
			NetSpeed string `yaml:"netSpeed"`
		} `yaml:"hostFact"`
		Rack string `yaml:"rack"`
	} `yaml:"hosts"`
	Services []struct {
		Type           string      `yaml:"type"`
		ID             string      `yaml:"id"`
		Name           interface{} `yaml:"name"`
		Description    string      `yaml:"description"`
		Mode           string      `yaml:"mode"`
		Version        string      `yaml:"version"`
		Releaseversion string      `yaml:"releaseversion"`
		ServiceHome    string      `yaml:"serviceHome"`
		ServiceConf    string      `yaml:"serviceConf"`
		Status         string      `yaml:"status"`
		IfAvailable    bool        `yaml:"ifAvailable"`
		Health         string      `yaml:"health"`
		Dependencies   []string    `yaml:"dependencies"`
		Roles          []struct {
			Type      string `yaml:"type"`
			RoleID    string `yaml:"roleId"`
			Instances []struct {
				ID       string `yaml:"id"`
				Idx      int    `yaml:"idx"`
				Hostname string `yaml:"hostname"`
				Status   string `yaml:"status"`
			} `yaml:"instances"`
			HostnameByRoles []string `yaml:"hostnameByRoles"`
		} `yaml:"roles"`
		Vars []struct {
			Name         string      `yaml:"name"`
			VarKey       interface{} `yaml:"varKey"`
			Type         string      `yaml:"type"`
			DefaultValue string      `yaml:"defaultValue"`
			Pattern      string      `yaml:"pattern"`
			Description  string      `yaml:"description"`
			ItemType     interface{} `yaml:"itemType"`
			ItemPattern  interface{} `yaml:"itemPattern"`
			Tags         []string    `yaml:"tags"`
		} `yaml:"vars"`
		Security struct {
			Type         string        `yaml:"type"`
			Dependencies []interface{} `yaml:"dependencies"`
		} `yaml:"security"`
		Metrics struct {
			YarnAllocatedMB          float64 `yaml:"yarn_AllocatedMB"`
			YarnGcCount              float64 `yaml:"yarn_GcCount"`
			YarnAllocatedVCores      float64 `yaml:"yarn_AllocatedVCores"`
			HdfsPercentBlockPoolUsed float64 `yaml:"hdfs_PercentBlockPoolUsed"`
			YarnAppsFailed           float64 `yaml:"yarn_AppsFailed"`
			YarnAvailableMB          float64 `yaml:"yarn_AvailableMB"`
			HdfsCorruptBlocks        float64 `yaml:"hdfs_CorruptBlocks"`
			HdfsMemHeapUsedM         float64 `yaml:"hdfs_MemHeapUsedM"`
			HdfsTotal                int64   `yaml:"hdfs_Total"`
			YarnGcTimeMillis         float64 `yaml:"yarn_GcTimeMillis"`
			YarnAppsKilled           float64 `yaml:"yarn_AppsKilled"`
			HdfsUsed                 int64   `yaml:"hdfs_Used"`
			HdfsNNStarted            string  `yaml:"hdfs_NNStarted"`
			YarnAppsPending          float64 `yaml:"yarn_AppsPending"`
			YarnMemHeapCommittedM    float64 `yaml:"yarn_MemHeapCommittedM"`
			HdfsFilesTotal           float64 `yaml:"hdfs_FilesTotal"`
			HdfsBlocksTotal          float64 `yaml:"hdfs_BlocksTotal"`
			YarnAppsRunning          float64 `yaml:"yarn_AppsRunning"`
			YarnLogFatal             float64 `yaml:"yarn_LogFatal"`
			HdfsClusterID            string  `yaml:"hdfs_ClusterId"`
			YarnAllocatedContainers  float64 `yaml:"yarn_AllocatedContainers"`
			HdfsMissingBlocks        float64 `yaml:"hdfs_MissingBlocks"`
			YarnLogError             float64 `yaml:"yarn_LogError"`
			YarnLogInfo              int     `yaml:"yarn_LogInfo"`
			YarnLogWarn              float64 `yaml:"yarn_LogWarn"`
			HdfsMemHeapCommittedM    float64 `yaml:"hdfs_MemHeapCommittedM"`
			HdfsSafemode             string  `yaml:"hdfs_Safemode"`
			YarnMemHeapMaxM          float64 `yaml:"yarn_MemHeapMaxM"`
			YarnMemHeapUsedM         float64 `yaml:"yarn_MemHeapUsedM"`
			YarnAvailableVCores      float64 `yaml:"yarn_AvailableVCores"`
			HdfsNumLiveDataNodes     float64 `yaml:"hdfs_NumLiveDataNodes"`
			HdfsPercentUsed          float64 `yaml:"hdfs_PercentUsed"`
		} `yaml:"metrics,omitempty"`
		Owner             string   `yaml:"owner"`
		HostnameByService []string `yaml:"hostnameByService"`
	} `yaml:"services"`
	CreateTime      string   `yaml:"createTime"`
	AllServicesType []string `yaml:"allServicesType"`
	AllHostNames    []string `yaml:"allHostNames"`
}

type ElasticsearchStruct struct {
	ClusterName                 string  `json:"cluster_name"`
	Status                      string  `json:"status"`
	TimedOut                    bool    `json:"timed_out"`
	NumberOfNodes               int     `json:"number_of_nodes"`
	NumberOfDataNodes           int     `json:"number_of_data_nodes"`
	ActivePrimaryShards         int     `json:"active_primary_shards"`
	ActiveShards                int     `json:"active_shards"`
	RelocatingShards            int     `json:"relocating_shards"`
	InitializingShards          int     `json:"initializing_shards"`
	UnassignedShards            int     `json:"unassigned_shards"`
	DelayedUnassignedShards     int     `json:"delayed_unassigned_shards"`
	NumberOfPendingTasks        int     `json:"number_of_pending_tasks"`
	NumberOfInFlightFetch       int     `json:"number_of_in_flight_fetch"`
	TaskMaxWaitingInQueueMillis int     `json:"task_max_waiting_in_queue_millis"`
	ActiveShardsPercentAsNumber float64 `json:"active_shards_percent_as_number"`
}
