import { DataSourceInstanceSettings } from '@kengine/data';
import { DataSourceWithBackend } from '@kengine/runtime';
import { MyDataSourceOptions, MyQuery } from './types';

export class DataSource extends DataSourceWithBackend<MyQuery, MyDataSourceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<MyDataSourceOptions>) {
    super(instanceSettings);
  }
}
