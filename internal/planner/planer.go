package planner

import (
	"fmt"
	"forge/internal/task"
	"sort"
)

func BuildGraph(conf *task.Config) map[string]*task.TaskNode {
	graph := make(map[string]*task.TaskNode)

	for name, t := range conf.Tasks {
		graph[name] = &task.TaskNode{
			Name:      name,
			Command:   t.Command,
			Args:      t.Args,
			DependsOn: t.DependsOn,
		}
	}

	return graph
}

func ResolveExecutionPlan(
	graph map[string]*task.TaskNode,
	target string,
) ([]*task.TaskNode, error) {

	node, ok := graph[target]
	if !ok {
		return nil, fmt.Errorf("target not found: %s", target)
	}

	visited := make(map[string]bool)
	visiting := make(map[string]bool)

	var ordered []*task.TaskNode

	if err := dfs(node, graph, visited, visiting, &ordered); err != nil {
		return nil, err
	}

	return ordered, nil
}

func dfs(
	node *task.TaskNode,
	graph map[string]*task.TaskNode,
	visited map[string]bool,
	visiting map[string]bool,
	ordered *[]*task.TaskNode,
) error {

	if visited[node.Name] {
		return nil
	}

	if visiting[node.Name] {
		return fmt.Errorf("cycle detected at task: %s", node.Name)
	}

	visiting[node.Name] = true

	dependencies := make([]string, len(node.DependsOn))
	copy(dependencies, node.DependsOn)

	sort.Strings(dependencies)

	for _, depName := range dependencies {
		depNode, ok := graph[depName]
		if !ok {
			return fmt.Errorf(
				"task [%s] depends on unknown task [%s]",
				node.Name,
				depName,
			)
		}

		if err := dfs(depNode, graph, visited, visiting, ordered); err != nil {
			return err
		}
	}

	visiting[node.Name] = false
	visited[node.Name] = true

	*ordered = append(*ordered, node)

	return nil
}
