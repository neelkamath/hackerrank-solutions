students = []
for _ in range(int(input())):
    students.append([input(), float(input())])
lowest = min([student[1] for student in students])
students = [student for student in students if student[1] != lowest]
second_lowest = min([student[1] for student in students])
names = [student[0] for student in students if student[1] == second_lowest]
names.sort()
print('\n'.join(names))
