
func dfs(cur, tgt Node, visited map[Node]struct{}) bool {
	if _, ex := visited[cur]; ex {
		return true
	}

	for next := range cur.Neighbor {
		if _, ex := visited[next]; !ex {
			visited[next] = struct{}{}
			if dfs(next, tgt, visited) {
				return true
			}
		}
	}

	return false
}

func dfs2(cur, tgt Node) bool {
}
