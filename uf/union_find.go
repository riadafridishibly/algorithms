package uf

type UnionFind struct {
	id    []int
	count int
}

func NewUnionFind(n int) *UnionFind {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	return &UnionFind{id: id, count: n}
}

func (u *UnionFind) Count() int {
	return u.count
}

func (u *UnionFind) Find(p int) int {
	return u.id[p]
}

func (u *UnionFind) Union(p, q int) {
	pID := u.Find(p)
	qID := u.Find(q)

	if pID == qID {
		return
	}
	for i := range u.id {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}

	u.count--
}

func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
