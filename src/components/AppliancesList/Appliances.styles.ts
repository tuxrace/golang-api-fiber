import { Theme } from "@material-ui/core";

export default (theme: Theme) => ({
   card: {
       [theme.breakpoints.down('xs')]: {
           minHeight: 220,
           margin: theme.spacing(0, 2)
       }
   }
})