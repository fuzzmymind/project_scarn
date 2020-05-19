<template>
  <div>
    <Converters />
    <br />
    <div class="container">
      <div class="card">
        <div class="card-body">
          <p>What is epoch time?</p>
          <p>
            The Unix epoch (or Unix time or POSIX time or Unix timestamp) is the
            number of seconds that have elapsed since January 1, 1970 (midnight
            UTC/GMT), not counting leap seconds (in ISO 8601:
            1970-01-01T00:00:00Z). Literally speaking the epoch is Unix time 0
            (midnight 1/1/1970), but 'epoch' is often used as a synonym for Unix
            time. Some systems store epoch dates as a signed 32-bit integer,
            which might cause problems on January 19, 2038 (known as the Year
            2038 problem or Y2038).
          </p>
        </div>
      </div>
      <br />
      <div>
        <table class="table">
          <thead>
            <tr>
              <th>Epoch To Human Readable</th>
              <th>Human Readable To Epoch</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <form
                  class="form-inline"
                  v-on:submit.prevent="fetchSomething"
                  action="/action_page.php"
                >
                  <input
                    type="text"
                    class="form-control"
                    placeholder="Enter input value"
                    id="inputValue"
                    v-model="inputValue"
                  />
                  <button type="submit" class="btn btn-primary">Go</button>
                </form>
              </td>

              <td>Otto</td>
            </tr>
            <tr>
              <td>
                <p>{{ result }}</p>
              </td>
              <td></td>
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
      result: "",
      inputValue: ""
    };
  },
  methods: {
    async fetchSomething() {
      const ip = await this.$axios.$get("/api/epoch?param=" + this.inputValue);
      this.result = ip;
    }
  }
};
</script>
