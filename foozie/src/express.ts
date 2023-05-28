import * as express from 'express';
import { AppClusterService } from './cluster.service';

function main() {
	const app = express();
  const port = 3001;

  app.get('/', (req, res) => {
    res.send('Hello World!');
  });

  app.listen(port, () => {
    console.log(`Example app listening on port ${port}`);
  });
}

AppClusterService.clusterize(main);
