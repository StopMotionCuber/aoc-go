


from ortools.linear_solver import pywraplp


def main(filename: str):
    full_sum = 0
    for line in open(filename).read().strip().split("\n"):
        if line == "":
            continue
        splitted = line.split(" ")
        activations = splitted[1:-1]
        joltages = list(map(int, splitted[-1][1:-1].split(",")))
        for i in range(len(activations)):
            activations[i] = list(map(int, activations[i][1:-1].split(",")))
        full_sum += solveProblem(joltages, activations)
    print("Final sum:", full_sum)



def solveProblem(joltages: list[int], activations: list[list[int]]):
    solver = pywraplp.Solver.CreateSolver('SCIP')
    activation_vars = []
    max_activations = max(joltages)
    # Variables
    for i in range(len(activations)):
        activation_vars.append(solver.IntVar(0, max_activations, f'activation_{i}'))
    # Constraints
    for i in range(len(joltages)):
        constraint = solver.RowConstraint(joltages[i], joltages[i], f'constraint_{i}')
        for j in range(len(activations)):
            if i in activations[j]:
                constraint.SetCoefficient(activation_vars[j], 1)
    # Objective, minimize total activations
    objective = solver.Objective()
    for j in range(len(activations)):
        objective.SetCoefficient(activation_vars[j], 1)
    objective.SetMinimization()
    status = solver.Solve()
    if status == pywraplp.Solver.OPTIMAL:
        print('Solution:')
        for j in range(len(activations)):
            print(f'Activation {j}: {activation_vars[j].solution_value()}')
        print(f'Optimal objective value: {objective.Value()}')
    return objective.Value()

if __name__ == "__main__":
    main("2025/10/input-user.txt")
