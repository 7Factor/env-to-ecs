provider "aws" {
  version = "~> 1.51"

  assume_role {
    role_arn = "arn:aws:iam::${var.deploy_to_account}:role/7FContinuousDelivery"
  }
}

# Use data to safeguard against a non-existing cluster during
# the CI step.
data "aws_ecs_cluster" "cluster" {
  cluster_name = "dev-starter"
}

# We name these things in a repeatable way, take advantage of this.
data "aws_security_group" "access_sg" {
  name = "${data.aws_ecs_cluster.cluster.cluster_name}-ecs-access-sg"
}

module "ecr_task" {
  source = "github.com/7factor/terraform-ecs-http-task"

  # Where we want to deploy the thing
  vpc_id       = "vpc-097d970e68e428986"
  cluster_name = "${data.aws_ecs_cluster.cluster.cluster_name}"

  # Information about what we're deploying
  app_name         = "${var.app_name}"
  app_port         = "${var.app_port}"
  service_role_arn = "arn:aws:iam::${var.deploy_to_account}:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS"

  # Load balancers and health checking
  health_check_path = "/status"
  lb_cert_arn       = "arn:aws:acm:us-east-1:024124389053:certificate/e59abffd-ba4d-4456-aa29-b7576cd52a9d"
  lb_public_subnets = ["subnet-04d48ee6429982edb", "subnet-0ed11e0dd4da2cd32", "subnet-02d3469fb31aeefad"]
  cluster_lb_sg_id  = "${data.aws_security_group.access_sg.id}"

  # Let's start with a hard coded container definition. A zero as host port means
  # we want an ephemeral range of ports.
  container_definition = <<EOF
[
  {
    "image": "${var.ecr_uri}:${var.ecr_tag}",
    "name": "${var.app_name}-cnt",
    "portMappings": [
      {
        "containerPort": ${var.app_port},
        "hostPort": 0
      }
    ],
    "environment": [
      {
        "name": "PORT",
        "value": "${var.app_port}"
      }
    ]
  }
]
EOF
}

variable "app_name" {
  default     = "golang-starter"
  description = "Name of our app."
}

variable "app_port" {
  default     = 3999
  description = "Port our application runs on. Hard coded for now."
}

variable "deploy_to_account" {
  description = "The account to deploy into. Passed in from concourse."
}

variable "ecr_uri" {
  description = "URI to the repo for the image to pull and deploy. Passed into the container definition."
}

variable "ecr_tag" {
  description = "Tag of the container to pull. Passed in from concourse."
}

output "lb_hostname" {
  value = "${module.ecr_task.lb_hostname}"
}

output "deployed_container_tag" {
  value = "${var.ecr_tag}"
}