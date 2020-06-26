import React, { useEffect, useState } from "react";
import axios from "axios";
import { API_URL } from "../../config";
import { Grid, Typography } from "@material-ui/core";
import Search from "../Search";
import ApplianceCard from "../ApplianceCard";

const AppliancesList = () => {
  const [appliances, setAppliances] = useState([]);

  useEffect(() => {
    async function getAppliances() {
      const res = await axios.get(`${API_URL}/appliances`);
      setAppliances(res.data);
    }
    getAppliances();
  }, []);

  const handleSearch = async (category: string, search: string) => {
    if (!category && !search) {
      const res = await axios.get(`${API_URL}/appliances`);
      setAppliances(res.data);
      return;
    }
    const res = await axios.get(
      `${API_URL}/appliances-search?category=${category}&search=${search}`
    );
    setAppliances(res.data);
  };

  return (
    <Grid container direction="row" alignItems="center">
      <Grid item>
        <h1>{`📺`} Appliances</h1>
      </Grid>
      <Grid item xs={12} lg={12}>
        <Search handleSearch={handleSearch} />
      </Grid>
      <Grid item container direction="column">
        {appliances.length > 0 ? (
          appliances.map((appliance: any) => <ApplianceCard data={appliance} />)
        ) : (
          <div>No Records Found</div>
        )}
      </Grid>
    </Grid>
  );
};

export default AppliancesList;
