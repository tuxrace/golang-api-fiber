import { Theme } from "@material-ui/core";

export default (theme: Theme) => ({
  card: {
    [theme.breakpoints.down("xs")]: {
      minHeight: 220,
      boxShadow:
        "0px 5px 5px -3px rgba(0,0,0,0.2), 0px 8px 10px 1px rgba(0,0,0,0.14), 0px 3px 14px 2px rgba(0,0,0,0.12)",
    },
  },
});
