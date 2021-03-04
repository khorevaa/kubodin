package cmd

import (
	"github.com/urfave/cli/v2"
)

type someCommand struct {
}

func (c *someCommand) run(context *cli.Context) error {

	return nil
}

func (c *someCommand) Cmd() *cli.Command {

	cmd := &cli.Command{
		//Category:    "some_category",
		Name:        "some_command",
		Usage:       "Usage decription",
		Description: `Full usage description `,
		Action:      c.run,
		Flags:       []cli.Flag{
			//&cli.StringFlag{
			//	Destination: &c.cluster, Name: "cluster-id",
			//	Value: "", Usage: "cluster uuid for join new work server"},
			//&cli.StringFlag{
			//	Destination: &c.Name, Name: "name", Aliases: []string{"N"},
			//	Value: "", Usage: "work server name", EnvVars: []string{"SERVER_NAME"}, Required: true},
			//&cli.StringFlag{
			//	Destination: &c.AgentHost, Name: "host",
			//	Value: "", Usage: "work server agent host", EnvVars: []string{"SERVER_AGENT_HOST"}, Required: true},
			//&cli.IntFlag{
			//	Destination: &c.AgentPort, Name: "port",
			//	Value: 1540, Usage: "work server agent port", EnvVars: []string{"SERVER_AGENT_HOST"}, DefaultText: "1540"},
			//&cli.StringFlag{
			//	Destination: &c.PortRange, Name: "port-range",
			//	Value: "1560:1591", Usage: "work server port range", EnvVars: []string{"SERVER_PORT_RANGE"}, DefaultText: "1560:1591"},
			//&cli.StringFlag{
			//	Destination: &c.Using, Name: "using",
			//	Value: "normal", Usage: "variant of using work server (main, normal)", EnvVars: []string{"SERVER_USING"}, DefaultText: "normal"},
			//&cli.StringFlag{
			//	Destination: &c.DedicateManagers, Name: "dedicate-managers",
			//	Value: "none", Usage: "вариант размещения менеджеров сервисов (all, none)", EnvVars: []string{"SERVER_DEDICATE_MANAGERS"}, DefaultText: "none"},
			//&cli.IntFlag{
			//	Destination: &c.ClusterPort, Name: "cluster-port",
			//	Value: 1541, Usage: "номер порта главного менеджера кластера", EnvVars: []string{"SERVER_CLUSTER_PORT"}, DefaultText: "1541"},
			//&cli.Int64Flag{
			//	Destination: &c.MemoryLimit, Name: "memory-limit",
			//	Value: 0, Usage: "предел использования памяти рабочими процессами (kilobytes)", EnvVars: []string{"SERVER_MEMORY_LIMIT"}},
			//&cli.Int64Flag{
			//	Destination: &c.ConnectionsLimit, Name: "connections-limit",
			//	Value: 128, Usage: "максимальное количество соединения на рабочий процесс", EnvVars: []string{"SERVER_CONNECTIONS_LIMIT"}, DefaultText: "128"},
			//&cli.Int64Flag{
			//	Destination: &c.CriticalTotalMemory, Name: "total-memory",
			//	Value: 0, Usage: "максимальный объем памяти процессов рабочего сервера (bytes)", EnvVars: []string{"SERVER_TOTAL_MEMORY"}},
		},
	}

	return cmd
}
