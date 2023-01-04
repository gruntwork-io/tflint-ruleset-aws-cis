# aws_security_group_rule_invalid_cidr_block

Disallow rules that allow `0.0.0.0/0` or `::/0` access on remote access control ports (22 and 3389).

## Example

```hcl
resource "aws_security_group_rule" "rule" {
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  type              = "ingress"
  cidr_blocks       = ["0.0.0.0/0", "10.0.0.0/16"]
}
```

```
1 issue(s) found:

Error: cidr_blocks can not contain '0.0.0.0/0' when allowing 'ingress' access to ports [22 3389] (aws_security_group_rule_invalid_cidr_block)

```

## Why

CIS AWS Benckmark has two recommendations regarding Security Group's CIDR blocks:
- 5.2 ensures no security groups allow ingress from 0.0.0.0/0 to remote server administration ports
- 5.3 ensures no security groups allow ingress from ::/0 to remote server administration ports

## How To Fix

Update `cidr_blocks` and/or `ipv6_cidr_blocks` to not allow access to the remote access ports, or update the port values to not contain the remote access ones.
