# amazonian
Opensource tool to faciliate AWS application resource development

## setup
Follow the setup process defined in the [AWS docs](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)

### Example command
`./workdir/amazonian --VPC=vpc-c7aa77be --Priority=12 --HostedZoneName=vssdevelopment.com --Image=willejs/go-hello-world --StackName=secondstack --ServiceName=Node --ContainerName=Hello --ClusterName=amazonian-ecs --ClusterExists`