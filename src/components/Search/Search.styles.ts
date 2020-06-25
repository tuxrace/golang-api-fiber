import { Theme } from "@material-ui/core";

export default (theme: Theme) => ({
    search: {
        marginBottom: theme.spacing(1)    
    },
    clear: {
        cursor: 'pointer',
        transition: '0.5s',
        opacity: '0.6',
        '&:hover': {
            opacity: '1'
        }
    }
})