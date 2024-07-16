import { VNExpress_selectorClient } from './Proto/VNExpress_selector_grpc_web_pb';
import { Range } from './Proto/VNExpress_selector_pb';

const client = new VNExpress_selectorClient('http://localhost:8080');


export const selectData = (type, mainCategories, subCategories, author, wholeDay, dayComparisor, releaseDay, timeComparisor, releaseTime, limit) => {
  const req = new Range();
  req.setMainCategoriesList(mainCategories);
  req.setSubCategoriesList(subCategories);
  req.setAuthorList(author);
  req.setWholeDay(wholeDay);
  req.setDayComparisorList(dayComparisor);
  req.setDayList(releaseDay);
  req.setTimeComparisor(timeComparisor);
  req.setTimeList(releaseTime);
  req.setLimit(limit);

  return new Promise((resolve, reject) => {
    const results = [];
    const stream = type === 'news' ? client.select_news(req, {}) : client.select_podcast(req, {});

    stream.on('data', (response) => {
      results.push(response.toObject());
    });

    stream.on('end', () => {
      resolve(results);
    });

    stream.on('error', (err) => {
      reject(err);
    });
  });
};
