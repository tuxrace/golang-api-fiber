import React from "react";
import {
  TextField,
  Grid,
  makeStyles,
  Button,
  FormControl,
  RadioGroup,
  FormLabel,
  FormControlLabel,
  Radio,
} from "@material-ui/core";
import styles from "./Search.styles";
import axios from "axios";
import { API_URL } from "../../config";

const useStyles = makeStyles(styles);

interface Props {
    handleSearch: (c: string, s: string) => void;
}

const Search: React.FC<Props> = ({handleSearch}) => {
  const classes = useStyles();
  const [category, setCategory] = React.useState('brand');
  const [search, setSearch] = React.useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setCategory((event.target as HTMLInputElement).value);
  };

  const submitSearch = () => {
      handleSearch(category, search);
  }

  const onSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const searchText = e.target.value;
    setSearch(searchText);
  }

  return (
    <Grid container>
      <Grid item xs={12} md={5}>
        <TextField
          type="text"
          color="secondary"
          placeholder="Search Appliance"
          fullWidth
          variant="filled"
          className={classes.search}
          value={search}
          onChange={onSearchChange}
        />
      </Grid>
      <Grid item xs={12} md={5} container justify="center">
        <FormControl>
          <RadioGroup name="search" value={category} onChange={handleChange} row>
            <FormControlLabel value="brand" control={<Radio />} label="Brand" />
            <FormControlLabel value="model" control={<Radio />} label="Model" />
            <FormControlLabel
              value="serialNumber"
              control={<Radio />}
              label="Serial Number"
            />
          </RadioGroup>
        </FormControl>
      </Grid>
      <Grid item container xs={12} md={2}>
        <Button color="primary" variant="contained" fullWidth  onClick={submitSearch} disabled={!search}>
          Search
        </Button>
      </Grid>
    </Grid>
  );
};

export default Search;
