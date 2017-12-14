package backend

import "github.com/sirupsen/logrus"

// this is main loop
func (b *backend) runDispatcher(events *SystemEvents) {
	for {
		select {
		case evt := <-events.Join:
			logrus.Debugf("[mainloop:Join] %+v", evt)
			b.userAgentManager.DeliverJoin(evt.UserName, evt.ChannelID)
		case evt := <-events.Leave:
			logrus.Debugf("[mainloop:Leave] %+v", evt)
			b.userAgentManager.DeliverLeave(evt.UserName, evt.ChannelID)
		case evt := <-events.CreateChannel:
			logrus.Debugf("[mainloop:CreateChannel] %+v", evt)
			b.channels[evt.Channel.ID] = evt.Channel
		case evt := <-events.UpdateChannel:
			logrus.Debugf("[mainloop:UpdateChannel] %+v", evt)
			// skip
		case evt := <-events.DeleteChannel:
			logrus.Debugf("[mainloop:DeleteChannel] %+v", evt)
			// skip
		case evt := <-events.ReceiveMessage:
			logrus.Debugf("[mainloop:ReceiveMessage] %s %s - %s", evt.ChannelID, evt.Message.Sender, evt.Message.Text)

			//b.channels[evt.ChannelID].AppendMessage(evt.Message)

			// TODO loop and find user

			for _, username := range b.channels[evt.ChannelID].Member {
				b.userAgentManager.DeliverMessage(username, evt.ChannelID, evt.Message)
				//or
				//go b.userAgentManager.DeliverMessage(username, evt.ChannelID, evt.Message)
			}

		case err := <-events.Error:
			// just leave log...
			logrus.Errorf("[mainloop:Error] %s", err.Error())
		}
	}
}
