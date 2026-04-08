@AI_CONTEXT.md

## Adding a new command

1. Create `src/app/commands/my_cmd.py` with `handle(args: str) -> None`
2. Register in `src/app/core.py`: `router.register("my")(my_cmd.handle)`
3. Add tests in `tests/test_commands.py`
