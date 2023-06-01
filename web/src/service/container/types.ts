export interface termResize {
  width: String
  height: String
}

export interface deleteContainer {
  host: string
  containerId: string[]
}

export interface Port {
  PrivatePort: number
  Type: string
}

export interface Labels {
  'com.docker.stack.namespace': string
  'com.docker.swarm.node.id': string
  'com.docker.swarm.service.id': string
  'com.docker.swarm.service.name': string
  'com.docker.swarm.task': string
  'com.docker.swarm.task.id': string
  'com.docker.swarm.task.name': string
}

export interface HostConfig {
  NetworkMode: string
}

export interface Bridge {
  IPAMConfig?: any
  Links?: any
  Aliases?: any
  NetworkID: string
  EndpointID: string
  Gateway: string
  IPAddress: string
  IPPrefixLen: number
  IPv6Gateway: string
  GlobalIPv6Address: string
  GlobalIPv6PrefixLen: number
  MacAddress: string
  DriverOpts?: any
}

export interface IPAMConfig {
  IPv4Address: string
}

export interface Ingress {
  IPAMConfig: IPAMConfig
  Links?: any
  Aliases?: any
  NetworkID: string
  EndpointID: string
  Gateway: string
  IPAddress: string
  IPPrefixLen: number
  IPv6Gateway: string
  GlobalIPv6Address: string
  GlobalIPv6PrefixLen: number
  MacAddress: string
  DriverOpts?: any
}

export interface IPAMConfig2 {
  IPv4Address: string
}

export interface PortainerAgentNetwork {
  IPAMConfig: IPAMConfig2
  Links?: any
  Aliases?: any
  NetworkID: string
  EndpointID: string
  Gateway: string
  IPAddress: string
  IPPrefixLen: number
  IPv6Gateway: string
  GlobalIPv6Address: string
  GlobalIPv6PrefixLen: number
  MacAddress: string
  DriverOpts?: any
}

export interface IPAMConfig3 {
  IPv4Address: string
}

export interface Macrowing {
  IPAMConfig: IPAMConfig3
  Links?: any
  Aliases?: any
  NetworkID: string
  EndpointID: string
  Gateway: string
  IPAddress: string
  IPPrefixLen: number
  IPv6Gateway: string
  GlobalIPv6Address: string
  GlobalIPv6PrefixLen: number
  MacAddress: string
  DriverOpts?: any
}

export interface Networks {
  bridge: Bridge
  ingress: Ingress
  portainer_agent_network: PortainerAgentNetwork
  macrowing: Macrowing
}

export interface NetworkSettings {
  Networks: Networks
}

export interface Mount {
  Type: string
  Name: string
  Source: string
  Destination: string
  Driver: string
  Mode: string
  RW: boolean
  Propagation: string
}

export interface Container {
  Id: string
  Names: string[]
  Image: string
  ImageID: string
  Command: string
  Created: number
  Ports: Port[]
  Labels: Labels
  State: string
  Status: string
  HostConfig: HostConfig
  NetworkSettings: NetworkSettings
  Mounts: Mount[]
  host: string
}
