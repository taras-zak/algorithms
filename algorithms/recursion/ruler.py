def draw_line(tick_len, label=""):
    line = "-" * tick_len
    if label:
        line += ' ' + label
    print(line)

def draw_interval(center_length):
    if center_length > 0:
        draw_interval(center_length - 1)
        draw_line(center_length)
        draw_interval(center_length - 1)

def draw_ruler(num_inches, major_lenth):
    draw_line(major_lenth, "0")
    for j in range(1, num_inches+1):
        draw_interval(major_lenth - 1)
        draw_line(major_lenth, str(j))


draw_ruler(1, 3)