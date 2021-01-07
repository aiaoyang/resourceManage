package common

import "errors"

var (
	errECSTransferError = errors.New("ecs response 类型不为 DescribeInstancesResponse")

	errDomainTransferError = errors.New("domain response 类型不为 QueryDomainListResponse")

	errCertTransferError = errors.New("cert response 类型不为 CommonResponse")

	errRDSTransferError = errors.New("rds response 类型不为 MyDescribeDBInstancesResponse")
)
