package ebs

type EBSVolume struct {
  Active bool
  Name string
  Tags string[string]
}

type EBS struct {
  Client string
  EBSVolumes []EBSVolume
}

func (e *ebs) InitClient() error {
  return nil
}

func (e *ebs) Volumes() error {
  return nil
}
