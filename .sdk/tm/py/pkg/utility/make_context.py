# BudDrive SDK utility: make_context

from projectname_sdk.core.context import BudDriveContext


def make_context_util(ctxmap, basectx):
    return BudDriveContext(ctxmap, basectx)
