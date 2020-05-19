<template>
  <div>
    <Converters />
    <br />
    <div class="container">
      <div class="card">
        <div class="card-body">
          <p>What is base64?</p>
          <p>Description goes here.</p>
        </div>
      </div>
      <br />
      <div>
        <table class="table">
          <thead>
            <tr>
              <th>Encode</th>
              <th>Decode</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <form class="form-inline" v-on:submit.prevent="encode" action="/action_page.php">
                  <input
                    type="text"
                    class="form-control"
                    placeholder="Enter input value"
                    id="inputEncValue"
                    v-model="inputEncValue"
                  />
                  <button type="submit" class="btn btn-primary">Go</button>
                </form>
              </td>

              <td>
                <form class="form-inline" v-on:submit.prevent="decode" action="/action_page.php">
                  <input
                    type="text"
                    class="form-control"
                    placeholder="Enter input value"
                    id="inputDecValue"
                    v-model="inputDecValue"
                  />
                  <button type="submit" class="btn btn-primary">Go</button>
                </form>
              </td>
            </tr>
            <tr>
              <td>
                <p>{{ resultEncode }}</p>
              </td>
              <td>
                <p>{{ resultDecode }}</p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Converters from "@/components/Navigation/Converters";
export default {
  components: {
    Converters
  },
  data() {
    return {
      resultEncode: "",
      resultDecode: "",
      inputEncValue: "",
      inputDecValue: ""
    };
  },
  methods: {
    async encode() {
      const ip = await this.$axios.$get(
        "/api/base64encode?param=" + this.inputEncValue
      );
      this.resultEncode = ip;
    },
    async decode() {
      const ip = await this.$axios.$get(
        "/api/base64decode?param=" + this.inputDecValue
      );
      this.resultDecode = ip;
    }
  }
};
</script>
