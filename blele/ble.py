pdb.runcall = debug


def debug(e, i, tb):
    traceback.print_exception(e, i, tb)
    pdb.pm

