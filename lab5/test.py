import random


def find(fn, itr):
    return next((x for x in itr if fn(x)), None)


def neighbours(src, E):
    return [q for (p, q) in E if p == src] + [p for (p, q) in E if q == src]


def proces(p, E, pref):
    N = neighbours(p, E)
    # accept proposal
    if pref[p] == None:
        q = find(lambda q: pref[q] == p, N)
        if q != None:
            pref[p] = q
    # propose
    if pref[p] == None and all([pref[q] != p for q in N]):
        q = find(lambda q: pref[q] == None, N)
        if q != None:
            pref[p] = q
    # unchain
    q = pref[p]
    if q != None and pref[q] != p and pref[q] != None:
        pref[p] = None


def married(p, E, pref):
    Np = neighbours(p, E)
    q = pref[p]
    Nq = neighbours(q, E)
    return (q in Np) and (pref[q] in Nq)


def single(p, E, pref):
    Np = neighbours(p, E)
    return pref[p] == None and all(married(q, E, pref) for q in Np)


def check_stabilization(V, E, pref):
    return all(married(p, E, pref) or single(p, E, pref) for p in V)


def build_mm(V, pref):
    M = [(p, pref[p]) for p in V if pref[p] != None]
    return [(i, j) for (i, j) in M if i < j]


def generate_star_graph(V):
    return [(V[0], j) for j in V if j != V[0]]


def main():
    random.seed(0)
    n = 11
    V = [i for i in range(n)]
    # E = [(i, j) for i in V for j in V if (i < j)]
    E = generate_star_graph(V) + [(2, 3), (4, 3)]
    print(E)

    # normalize
    E = list(map(lambda e: e if e[0] < e[1] else (e[1], e[0]), E))
    print(E)

    pref=dict()
    for p in V:
        pref[p]=None

    stabilized=False
    while not stabilized:
        for p in V:
            proces(p, E, pref)
            stabilized=check_stabilization(V, E, pref)
            if stabilized:
                break

    M=build_mm(V, pref)
    print(M)



if __name__ == "__main__":
    main()
