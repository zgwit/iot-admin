import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {DashComponent} from './dash/dash.component';
import {MainComponent} from './main.component';
import {PasswordComponent} from './password/password.component';
import {LinkComponent} from './link/link.component';
import {ChannelComponent} from './channel/channel.component';
import {UserComponent} from './user/user.component';

const routes: Routes = [
  {
    path: '',
    component: MainComponent,
    children: [
      {path: '', redirectTo: 'dash'},
      {path: 'dash', component: DashComponent},
      {path: 'channel', component: ChannelComponent},
      {path: 'link', component: LinkComponent},
      {path: 'user', component: UserComponent},
      {path: 'password', component: PasswordComponent},
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MainRoutingModule {
}